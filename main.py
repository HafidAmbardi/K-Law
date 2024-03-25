import os
from dotenv import find_dotenv, load_dotenv
import openai 
import time
import logging
from datetime import datetime


load_dotenv()

openai.api_key = os.environ.get("OPENAI_API_KEY")

client = openai.OpenAI()
model = "gpt-3.5-turbo-16k"

# Creating our Assistant
k_law_assistant = client.beta.assistants.create(
    name="K-Law",
    instructions="""You are playing the role of a Korean Labor Law Expert.  A user will provide a written description of a workplace situation.  Based on this description, do your best to:\n

1. Summarize the key points of the scenario.\n
2. Analyze the scenario to identify potential violations of Korean Labor Law.\n
3. If a violation is found, pinpoint the relevant Korean Labor Law acts, chapters, sections, and articles.\n
4. Recommend potential next steps for the user, emphasizing that they should consult with a qualified lawyer for specific legal advice.\n""",
model=model,
)

assistant_id = k_law_assistant.id
print(k_law_assistant.id)


# Thread
# thread = client.beta.threads.create(
#     messages=[
#         {
#             "role": "user",
#             "content":"My employer has asked me to work 60 hours a week, but I am only being paid for 40 hours. I am worried that this is illegal. What should I do?"
#         }
#     ]
# )

# thread_id = thread
# print(thread_id)

# Hardcode our ids
assistant_id="asst_PXPtai3NuiUDHSHEwh4n3Zzz"
thread_id = "thread_eFz2UMZx9v78B2mB2le35WMI"

# Create message
message = "My employer has asked me to work 60 hours a week, but I am only being paid for 40 hours. I am worried that this is illegal. What should I do?"
message = client.beta.threads.messages.create(
    thread_id=thread_id,
    role="user",
    content=message,
)

# Run our assistant
run = client.beta.threads.runs.create(
    thread_id=thread_id,
    assistant_id=assistant_id,
    instructions="Please talk like you're a chicken.",
)