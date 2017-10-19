import twillio

def two_factor(context):
  # Generate a random six digit code
  code = random.randint(100000, 999999)
  
  # Register the code with the login service
  user = context.json["user"]
  register_code_with_login_service(user, code)
  
  # Use the twillio library to send texts
  account = "my-account-sid"
  token = "my-token"
  client = twilio.rest.Client(account, token)

  user_number = context.json["phoneNumber"]
  msg = "Hello {} your authentication code is: {}.".format(user, code)
  message = client.api.account.messages.create(to=user_number,
                                               from_="+12065251212",
                                               body=msg)
  return {"status": "ok"}
