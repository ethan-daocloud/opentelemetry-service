// Copyright 2019, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Program otelsvc is the OpenTelemetry Collector that collects stats
// and traces and exports to a configured backend.
package main

import (
	"log"

	"github.com/open-telemetry/opentelemetry-service/defaults"
	"github.com/open-telemetry/opentelemetry-service/service"
)

func main() {
	handleErr := func(err error) {
		if err != nil {
			log.Fatalf("Failed to run the service: %v", err)
		}
	}

	factories, err := defaults.Components()
	handleErr(err)

	svc := service.New(factories)
	err = svc.StartUnified()
	handleErr(err)
}
