#!/usr/bin/env bash

. ./test/helpers.sh
. ./bin/check

setUp() {
  declare mkpcli_output=""
  declare mkpcli_return=0
  declare mkpcli_expected_args=""
  function mkpcli() {
    if [ ! "$*" = "${mkpcli_expected_args}" ]; then
      echo "ASSERT:expected:<${mkpcli_expected_args}> but was:<$*>" >&2
      return 1
    fi
    echo "${mkpcli_output}"
    return ${mkpcli_return}
  }
}

testItReturnsTheListOfVersions() {
  export CSP_API_TOKEN="my-token"
  export PRODUCT_SLUG="my-product"
  input=$(buildInput)
  mkpcli_expected_args="product list-versions --product my-product --output json"

  mkpcli_output='[{
    "versionnumber": "1.1.1",
    "extra_info": "not_needed"
  }, {
    "versionnumber": "2.2.2",
    "extra_info": "not_needed"
  }]'
  result=$(check "${input}" 2>&1)
  assertEquals "${?}" 0
  assertEquals "$(echo "${result}" | jq -r .[0].versionnumber)" "2.2.2"
  assertEquals "$(echo "${result}" | jq -r .[1].versionnumber)" "1.1.1"
}

testWhenCSIAPITokenIsEmptyItReturnsAnError() {
  export CSP_API_TOKEN=""
  export PRODUCT_SLUG="my-product"
  input=$(buildInput)

  result=$(check "${input}" 2>&1)
  assertEquals "${?}" 1
  assertContains "${result}" "Missing CSP API Token. Please ensure that source.csp_api_token has been set."
}

testWhenProductSlugIsEmptyItReturnsAnError() {
  export CSP_API_TOKEN="my-token"
  export PRODUCT_SLUG=""
  input=$(buildInput)

  result=$(check "${input}" 2>&1)
  assertEquals "${?}" 1
  assertContains "${result}" "Missing Marketplace product slug. Please ensure that source.product_slug has been set."
}

testWhenMkpcliReturnsAnErrorItReturnsAnError() {
  export CSP_API_TOKEN="my-token"
  export PRODUCT_SLUG="my-product"
  input=$(buildInput)
  mkpcli_expected_args="product list-versions --product my-product --output json"

  mkpcli_return=1
  result=$(check "${input}" 2>&1)
  assertEquals "${?}" 1
}

. ./test-lib/shunit2/shunit2
