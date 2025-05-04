#!/usr/bin/env bash
 
 # Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
 # Use of this source code is governed by a MIT style
 # license that can be found in the LICENSE file.
 
 # Common utilities, variables and checks for all build scripts.
 set -o errexit
 set -o nounset
 set -o pipefail
 
 # The root of the build/dist directory
 PROJ_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
 source ${PROJ_ROOT}/scripts/lib/logging.sh
 
 INSECURE_SERVER="127.0.0.1:8080"
 
 Header="-HContent-Type: application/json"
 CCURL="curl -f -s -XPOST" # Create
 UCURL="curl -f -s -XPUT" # Update
 RCURL="curl -f -s -XGET" # Retrieve
 DCURL="curl -f -s -XDELETE" # Delete
 
 mb::test::login()
 {
   ${CCURL} "${Header}" http://${INSECURE_SERVER}/login \
     -d'{"username":"root","password":"miniblog1234"}' | grep -Po 'token[" :]+\K[^"]+'
 }
 
 mb::test::user()
 {
   token="-HAuthorization: Bearer $(mb::test::login)"
 
   ${DCURL} "${token}" http://${INSECURE_SERVER}/v1/users/colin > /dev/null
   ${DCURL} "${token}" http://${INSECURE_SERVER}/v1/users/posttest > /dev/null
 
   ${CCURL} "${Header}" "${token}" http://${INSECURE_SERVER}/v1/users \
     -d'{"username":"colin","password":"miniblog1234","nickname":"belm","email":"nosbelm@qq.com","phone":"1818888xxxx"}' > /dev/null
 
   ${RCURL} "${token}" "http://${INSECURE_SERVER}/v1/users?offset=0&limit=10" > /dev/null
 
   ${RCURL} "${token}" http://${INSECURE_SERVER}/v1/users/colin > /dev/null
 
   ${UCURL} "${Header}" "${token}" http://${INSECURE_SERVER}/v1/users/colin \
     -d'{"nickname":"colin","email":"colin_modified@foxmail.com","phone":"1812884xxxx"}' > /dev/null
 
   ${DCURL} "${token}" http://${INSECURE_SERVER}/v1/users/colin > /dev/null
   mb::log::info "$(echo -e '\033[32mcongratulations, /v1/users test passed!\033[0m')"
 }
 
 mb::test::post()
 {
 
   ${CCURL} "${Header}" "${token}" http://${INSECURE_SERVER}/v1/users \
     -d'{"username":"posttest","password":"miniblog1234","nickname":"belm","email":"nosbelm@qq.com","phone":"1818888xxxx"}' > /dev/null
 
   tokenStr=`${CCURL} "${Header}" http://${INSECURE_SERVER}/login -d'{"username":"posttest","password":"miniblog1234"}' | grep -Po 'token[" :]+\K[^"]+'`
   token="-HAuthorization: Bearer ${tokenStr}"
 
   postID=`${CCURL} "${Header}" "${token}" http://${INSECURE_SERVER}/v1/posts -d'{"title":"installation","content":"installation."}' | grep -Po 'post-[a-z0-9]+'`
 
   ${RCURL} "${token}" http://${INSECURE_SERVER}/v1/posts > /dev/null
 
   ${RCURL} "${token}" http://${INSECURE_SERVER}/v1/posts/${postID} > /dev/null
 
   ${UCURL} "${Header}" "${token}" http://${INSECURE_SERVER}/v1/posts/${postID} -d'{"title":"modified"}' > /dev/null
 
   ${DCURL} "${token}" http://${INSECURE_SERVER}/v1/posts/${postID} > /dev/null
   mb::log::info "$(echo -e '\033[32mcongratulations, /v1/posts test passed!\033[0m')"
 }
 
 mb::test::user
 
 mb::test::post
 
 mb::log::info "$(echo -e '\033[32mcongratulations, all test passed!\033[0m')"