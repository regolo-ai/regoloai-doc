# Private Endpoints

Deploy private endpoints for enhanced security and control.

## Overview

Private Endpoints provide dedicated, isolated API endpoints running on your own infrastructure or Regolo's isolated environment. This ensures your API traffic never traverses the public internet.

## Benefits

- **Data Isolation**: Dedicated infrastructure with no shared resources
- **Network Security**: VPC/VPN integration options available
- **Compliance**: Meet regulatory requirements for data residency
- **Performance**: Dedicated bandwidth and processing capacity
- **Control**: Full control over access, authentication, and rate limiting

## Deployment Options

### Option 1: Regolo-Hosted Private Endpoint

Regolo manages the infrastructure on isolated, dedicated hardware:

```bash
Endpoint: https://private-<org-id>.regolo.ai/v1
Access: IP whitelisting, API key authentication
```

### Option 2: VPC Endpoint (AWS)

Integrate directly with your AWS VPC:

```bash
# AWS PrivateLink configuration
Service Name: com.amazonaws.regolo.vpce.us-east-1
```

## Setup

1. Contact enterprise support to enable Private Endpoints
2. Configure network access policies
3. Update your application endpoints
4. Test connectivity in staging environment
5. Deploy to production

## Authentication

Private endpoints support the same authentication methods as public endpoints:

=== "API Key"

    ```python
    import regolo
    
    regolo.default_key = "YOUR_API_KEY"
    regolo.base_url = "https://private-<org-id>.regolo.ai"
    ```

=== "OAuth 2.0"

    ```python
    from requests_oauthlib import OAuth2Session
    
    client = OAuth2Session(
        client_id,
        token=token,
        redirect_uri="https://private-<org-id>.regolo.ai/callback"
    )
    ```

## Monitoring

Monitor your private endpoint usage:

- Real-time request metrics
- Bandwidth monitoring
- Error rate tracking
- IP access logs