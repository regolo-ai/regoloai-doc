# Getting Started with Regolo.ai

To get started with Regolo.ai, sign up for an account at [dashboard.regolo.ai](https://dashboard.regolo.ai).

## Generate an API Key

Once logged in, navigate to the **Virtual Keys** section and create a new key. You can choose a specific model or select "All models" to use the key across all available models.

## Choose your client

Regolo.ai is fully compatible with the OpenAI API, so you can use either:

[**Regolo Python Library**](https://pypi.org/project/regolo/) or [**OpenAI Python Library**](https://pypi.org/project/openai/)

## Chat Example

Below is an example of how to create a simple chat application in python using regolo client.

=== "Using Regolo Client"

    ```python
        import streamlit as st
        import regolo
        
        regolo.default_key = "YOUR-API-KEY-HERE"
        regolo.default_chat_model = "Llama-3.3-70B-Instruct"
        
        client = regolo.RegoloClient()
        
        st.title("Regolo.ai Chat")
        
        if "messages" not in st.session_state:
            st.session_state.messages = []
        
        for msg in st.session_state.messages:
            with st.chat_message(msg["role"]):
                st.markdown(msg["content"])
        
        user_input = st.chat_input("Write a message...")
        if user_input:
            st.session_state.messages.append({"role": "user", "content": user_input})
            
            with st.chat_message("user"):
                st.markdown(user_input)
            
            client.add_prompt_to_chat(role="user", prompt=user_input)
            for msg in st.session_state.messages:
                client.add_prompt_to_chat(role=msg["role"], prompt=msg["content"])
            
            response = client.run_chat()
        
            role, content = response
            
            st.session_state.messages.append({"role": role, "content": content})
            
            with st.chat_message(role):
                st.markdown(content)
    ```
