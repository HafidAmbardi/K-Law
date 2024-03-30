import os
from dotenv import load_dotenv
import openai
import requests
import json

import time
import logging
from datetime import datetime
import streamlit as st


load_dotenv()

client = openai.OpenAI()

model = "gpt-4-1106-preview"  # "gpt-3.5-turbo-16k"

# Step 1. Upload a file to OpenaI embeddings ===
filepath = "./ENFORCEMENT_DECREE OF_THE_LABOR_STANDARDS_ACT.pdf"
file_object = client.files.create(file=open(filepath, "rb"), purpose="assistants")

# # Step 2 - Create an assistant
# assistant = client.beta.assistants.create(
#     name="KLaw",
#     instructions="""You are playing the role of a Korean Labor Law Expert.  A user will provide a written description of a workplace situation.  Based on this description, do your best to:
# 1. Summarize the key points of the scenario.
# 2. Analyze the scenario to identify potential violations of Korean Labor Law.
# 3. If a violation is found, retrieve the PDF data from your custom knowledge base to pinpoint the relevant Korean Labor Law acts, chapters, sections, and articles.
# 4. Recommend potential next steps for the user, emphasizing that they should consult with a qualified lawyer for specific legal advice.
# """,
#     tools=[{"type": "retrieval"}],
#     model=model,
#     file_ids=[file_object.id],
# )

# # === Get the Assis ID ===
# assis_id = assistant.id
# print(assis_id)

# # == Hardcoded ids to be used once the first code run is done and the assistant was created
thread_id = "thread_9OoWWGjU9WxAwT5KjSJyy42w"
assis_id = "asst_3qBgFUBM0wOERkyOeADjJJp7"

# == Step 3. Create a Thread
message = "What are Labor Laws in Korea?"

# thread = client.beta.threads.create()
# thread_id = thread.id
# print(thread_id)

message = client.beta.threads.messages.create(
    thread_id=thread_id, role="user", content=message
)

# == Run the Assistant
run = client.beta.threads.runs.create(
    thread_id=thread_id,
    assistant_id=assis_id,
)


def wait_for_run_completion(client, thread_id, run_id, sleep_interval=5):
    """
    Waits for a run to complete and prints the elapsed time.:param client: The OpenAI client object.
    :param thread_id: The ID of the thread.
    :param run_id: The ID of the run.
    :param sleep_interval: Time in seconds to wait between checks.
    """
    while True:
        try:
            run = client.beta.threads.runs.retrieve(thread_id=thread_id, run_id=run_id)
            if run.completed_at:
                elapsed_time = run.completed_at - run.created_at
                formatted_elapsed_time = time.strftime(
                    "%H:%M:%S", time.gmtime(elapsed_time)
                )
                print(f"Run completed in {formatted_elapsed_time}")
                logging.info(f"Run completed in {formatted_elapsed_time}")
                # Get messages here once Run is completed!
                messages = client.beta.threads.messages.list(thread_id=thread_id)
                last_message = messages.data[0]
                response = last_message.content[0].text.value
                print(f"Assistant Response: {response}")
                break
        except Exception as e:
            logging.error(f"An error occurred while retrieving the run: {e}")
            break
        logging.info("Waiting for run to complete...")
        time.sleep(sleep_interval)


# == Run it
wait_for_run_completion(client=client, thread_id=thread_id, run_id=run.id)

# === Check the Run Steps - LOGS ===
run_steps = client.beta.threads.runs.steps.list(thread_id=thread_id, run_id=run.id)
print(f"Run Steps --> {run_steps.data[0]}")
