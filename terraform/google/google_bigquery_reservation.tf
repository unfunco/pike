resource "google_bigquery_reservation" "pike" {
  name     = "my-reservation"
  location = "europe-west2"
  // Set to 0 for testing purposes
  // In reality this would be larger than zero
  slot_capacity     = 0
  edition           = "STANDARD"
  ignore_idle_slots = true
  concurrency       = 0
  autoscale {
    max_slots = 100
  }
}
