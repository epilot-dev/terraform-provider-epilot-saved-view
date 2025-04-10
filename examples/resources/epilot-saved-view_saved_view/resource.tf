resource "epilot-saved-view_saved_view" "my_savedview" {
  name   = "View listing German"
  shared = true
  slug = [
    "contact"
  ]
  ui_config = "{ \"see\": \"documentation\" }"
}