/*
 *Copyright (c) 2022, kaydxh
 *
 *Permission is hereby granted, free of charge, to any person obtaining a copy
 *of this software and associated documentation files (the "Software"), to deal
 *in the Software without restriction, including without limitation the rights
 *to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *copies of the Software, and to permit persons to whom the Software is
 *furnished to do so, subject to the following conditions:
 *
 *The above copyright notice and this permission notice shall be included in all
 *copies or substantial portions of the Software.
 *
 *THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 *SOFTWARE.
 */
package options

import (
	"context"

	opentelemetry_ "github.com/kaydxh/golang/pkg/opentelemetry"
	webserver_ "github.com/kaydxh/golang/pkg/webserver"
	"github.com/sirupsen/logrus"
)

// installOpenTelemetryOrDie initializes TracerProvider and MeterProvider
// Must be called after webserver.New() and before webserver.Run()
func (s *CompletedServerRunOptions) installOpenTelemetryOrDie(ctx context.Context, ws *webserver_.GenericWebServer) {
	// Set gin router for /metrics endpoint registration (Prometheus Pull mode)
	s.opentelemetryConfig.ApplyOptions(opentelemetry_.WithGinRouter(ws.GetGinEngine()))

	c := s.opentelemetryConfig.Complete()

	// Install Tracer
	err := c.InstallTracer(ctx)
	if err != nil {
		logrus.WithError(err).Fatalf("install OpenTelemetry Tracer, exit")
		return
	}

	// Install Meter
	err = c.InstallMeter(ctx)
	if err != nil {
		logrus.WithError(err).Fatalf("install OpenTelemetry Meter, exit")
		return
	}
}
