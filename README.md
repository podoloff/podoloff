# Podoloff
The Open Source Content Moderator

## Vision

The internet has given rise to many different communities. You can become friends with someone on the other side of the world in an instant. However, as these communities have grown, so has their complexity. Each community is different in the guidelines and standards they have for interactions. Current mediums of communication for these communities (Facebook groups, Slack channels, Github repos, etc.) unfortunately make it very difficult to moderate the content that is shared. Furthermore, developers who want to build new communities experience a large technical design barrier when thinking about how they want users to interact on their platform, and once a decision is made, it is diffult to dynamically update and modify those rules.

Podoloff exists to make content moderation easier with the following core tenants:
- **Open**: Podoloff is built in the open because it is important for content moderators to be able to see that the rules that they setup are what is governing content, not the Podoloff platform itself.
- **Pluggable**: Podoloff doesn't do anything without plugins. We don't make any assumptions about where communities should exist. If you want to use it for a Slack channel, there will be an plugin for that. If you want to use it for a Facebook group, there will be an plugin for that. If there is not an integration that exists for the platform you use, you can make one and use it with Podoloff. Lastly, if you are a developer building your own community, you will be able use the Podoloff SDK to quickly setup content moderation for your application.
- **Secure**: There will be a SAAS based Podoloff product that you can use if you want to (thank you to those of you that do!), but we make Podoloff so that you can easily run your own instance on your favorite cloud provider or your on-prem hardware. This puts you in control of the security of the system (we will also create opportunities to integrate with identity provider solutions like Active Directory and other LDAP-based services).

## Contributing

We intend Podoloff to be built to satisfy the needs of anyone who leads a community, so we welcome anyone to contribute to the project. If you think there is an area for improvement, take a look at how we do development in CONTRIBUTING.md.