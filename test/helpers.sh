#!/usr/bin/env bash

function buildInput() {
  jq -n \
    --arg cspAPIToken "${CSP_API_TOKEN}" \
    --arg productSlug "${PRODUCT_SLUG}" \
    '{
      "source": {
        "csp_api_token": $cspAPIToken,
        "product_slug": $productSlug
      }
    }'
}
