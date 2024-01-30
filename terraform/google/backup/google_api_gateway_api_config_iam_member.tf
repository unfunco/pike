resource "google_api_gateway_api_config_iam_member" "pike" {
  provider   = google-beta
  api        = google_api_gateway_api_config.pike.api
  api_config = google_api_gateway_api_config.pike.api_config_id
  role       = "roles/apigateway.viewer"
  member     = "user:crwoolfenden@gmail.com"
}
