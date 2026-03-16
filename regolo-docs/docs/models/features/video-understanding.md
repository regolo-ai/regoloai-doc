# Video Understanding

Video understanding enables the processing of video content alongside text, allowing for a wide range of applications such as video summarization, content analysis, action recognition, and information extraction from visual and audio content. By sending a combination of text prompts and video URLs, the model can provide insightful responses based on the video input.

!!! warning
    The video must be publicly accessible via a URL. Ensure the video URL is directly downloadable and does not require authentication or additional permissions.

## Video Completions

You can provide videos using a publicly accessible URL:

* [**Remote Video URL**](#remote-video-url): Supply a publicly accessible URL pointing to the video file.

### Remote Video URL

This section demonstrates how to use a remote video URL with the video understanding model.

=== "Using Regolo Client"
    ```python
    import regolo

    regolo.default_key = "YOUR_REGOLO_KEY"
    regolo.default_chat_model = "qwen3.5-122b"

    print(regolo.static_chat_completions(messages=[{
        "role": "user",
        "content": [
            {
                "type": "text",
                "text": "Please provide a detailed summary of this video."
            },
            {
                "type": "video_url",
                "video_url": {
                    "url": "https://your-public-video.url/video.mp4"
                }
            }
        ]
    }]))
    ```

=== "Python"
    ```python
    import requests

    url = "https://api.regolo.ai/v1/chat/completions"
    payload = {
        "model": "qwen3.5-122b",
        "messages": [
            {
                "role": "user",
                "content": [
                    {
                        "type": "text",
                        "text": "Please provide a detailed summary of this video."
                    },
                    {
                        "type": "video_url",
                        "video_url": {
                            "url": "https://your-public-video.url/video.mp4"
                        }
                    }
                ]
            }
        ]
    }
    headers = {
        "Content-Type": "application/json",
        "Authorization": "Bearer YOUR_REGOLO_KEY"
    }

    response = requests.post(url, json=payload, headers=headers)
    print(response.json())
    ```


=== "CURL"
    ```bash
    curl -X POST https://api.regolo.ai/v1/chat/completions \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer YOUR_REGOLO_KEY" \
    -d '{
        "model": "qwen3.5-122b",
        "messages": [
            {
                "role": "user",
                "content": [
                    {
                        "type": "text",
                        "text": "Provide a detailed summary of this video."
                    },
                    {
                        "type": "video_url",
                        "video_url": {
                            "url": "https://your-public-video.url/video.mp4"
                        }
                    }
                ]
            }
        ]
    }'
    ```
