def email_user(context):
  # Get the user name
  user = context.json['username']
  
  msg = 'Hello {} thanks for joining my awesome service!".format(user)
  
  send_email(msg, contex.json['email])


def subscribe_user(context):
  # Get the user name
  email = context.json['email']
  subscribe_user(email)