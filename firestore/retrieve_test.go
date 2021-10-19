// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package firestore

import (
	"context"
	"os"
	"reflect"
	"runtime"
	"testing"

	"cloud.google.com/go/firestore"
)

func TestRetrieve(t *testing.T) {
	// TODO(#559): revert this to testutil.SystemTest(t).ProjectID
	// when datastore and firestore can co-exist in a project.
	projectID := os.Getenv("GOLANG_SAMPLES_FIRESTORE_PROJECT")
	if projectID == "" {
		t.Skip("Skipping firestore test. Set GOLANG_SAMPLES_FIRESTORE_PROJECT.")
	}

	ctx := context.Background()

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		t.Fatalf("firestore.NewClient: %v", err)
	}
	defer client.Close()

	must := func(f func(context.Context, *firestore.Client) error) {
		err := f(ctx, client)
		if err != nil {
			fn := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
			t.Fatalf("%s: %v", fn, err)
		}
	}

	must(prepareRetrieve)

	_, err = docAsMap(ctx, client)
	if err != nil {
		t.Fatalf("Cannot get doc as map: %v", err)
	}

	_, err = docAsEntity(ctx, client)
	if err != nil {
		t.Fatalf("Cannot get doc as entity: %v", err)
	}

	must(multipleDocs)
	must(allDocs)
	must(getCollections)
}
