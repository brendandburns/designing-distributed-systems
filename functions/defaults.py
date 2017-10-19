# Simple handler function for adding default values
def handler(context):
  # Get the input value
  obj = context.json
  # If the 'name' field is not present, set it randomly
  if obj.get("name", None) is None:
    obj["name"] = random_name()
  # If the 'color' field is not present, set it to 'blue'
  if obj.get("color", None) is None:
    obj["color"] = "blue"
  # Call the actual API, potentially with the new default
  # values, and return the result
  return call_my_api(obj)
