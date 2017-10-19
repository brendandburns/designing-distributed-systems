def create_user(context):
  # For required event handlers, call them universally
  for key, value in required.items():
    call_function(value.webhook, context.json)
  
  # For optional event handlers, check and call them
  # conditionally
  for key, value in optional.items():
    if context.json.get(key, None) is not None:
      call_function(value.webhook, context.json)
