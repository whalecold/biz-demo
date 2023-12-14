// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

include "base.thrift"
namespace go cwg.bookinfo.ratings


struct RatingReq {
    1: required string ProductID,
}

struct RatingResp {
    1: required i8 Rating,
	2: required string RatingsInstance,
}

service RatingService {
    RatingResp ratings(1: RatingReq req)
}
