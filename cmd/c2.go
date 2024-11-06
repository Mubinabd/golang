package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

const (
	service     = "example-api"
	environment = "development"
	id          = 1
)

type Message struct {
	Text string `json:"text"`
}

func initTracer() (*sdktrace.TracerProvider, error) {
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint("http://localhost:14268/api/traces")))
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(service),
			attribute.String("environment", environment),
			attribute.Int64("ID", id),
		)),
	)

	otel.SetTracerProvider(tp)
	return tp, nil
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tracer := otel.Tracer(service)

	_, span := tracer.Start(ctx, "helloHandler")
	defer span.End()

	time.Sleep(50 * time.Millisecond)

	span.SetAttributes(attribute.String("handler", "hello"))

	msg := Message{Text: "Hello, World!"}
	json.NewEncoder(w).Encode(msg)
}

func main() {
	tp, err := initTracer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	handler := otelhttp.NewHandler(http.HandlerFunc(helloHandler), "hello")

	http.Handle("/hello", handler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
