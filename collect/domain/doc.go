// Copyright 2020 The Goava authors.
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

// Package domain contains descriptors for discrete comparable domains such as all int instances.
//
// A discrete domain is one that supports the three basic operations: Next, Previous and Distance, according to their
// specifications.
//
// A discrete domain always represents the entire set of values of its type; it cannot represent partial domains such as
// "prime integers" or "strings of length 5."
package domain
