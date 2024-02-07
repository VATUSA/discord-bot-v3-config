# discord-bot-v3-config

Configuration repository for [VATUSA/discord-bot-v3](https://github.com/VATUSA/discord-bot-v3)

All configuration changes should be made via a fork and be submitted via a Pull Request to main.

# Setup Steps
- Create a config file per the example with the appropriate values for your server (documentation can be found below)
- Create a Pull Request with your config file
- Add the VATUSA Integration bot to your discord server (link below)
- Ensure the bot's role is sorted above all roles that will need to be managed by the bot. The bot will not be able to assign roles or adjust nickname for any user who has a role higher than the bot's highest role.
- Post a link to your PR in #wm-chat in the VATUSA discord
- NOTE: The bot will not become active in your server until the configuration is approved and deployed. Please note on the PR if you need us to coordinate with you before activating the bot.

If you require assistance with any of the setup steps above, feel free to reach out in #discord-bot in the VATUSA discord.

# Recommendations / Considerations
- Consider adding a role for ACE Team, if you don't already have one. This will allow ACE Controllers to quickly join your discord and be able to access coordination rooms if needed for an event.
- Consider adding a #verification (or similarly named) channel to let people know how to link their discord account to get their roles. We recommend having a role for all 'verified' members and hiding this channel from all who are already verified. (example text below)
- Consider adding a role for neighboring controllers, if you do coordination in discord.

# Example verification prompt
If you can see this channel, your discord account isn't linked with your VATUSA account. Please visit https://www.vatusa.net/my/profile to link your discord account and receive your roles.

# Bot Add link
https://discord.com/api/oauth2/authorize?client_id=1122353511322103808&permissions=8&scope=bot


# Server Configuration
[Example configuration (ZSE)](https://github.com/VATUSA/discord-bot-v3/blob/main/config/servers/zse.yaml)

Required Configuration
```yaml
id: <The ID of the discord 'guild'>
name: <Short name of the server, used for logging>
active: <true/false>
description: <description of the purpose of the server - not used programmatically>
facility: <facility ID or ALL for division owned servers>
name_format_type: <how names should be formatted, see options below>
title_type: <what title should be appended to the name, see options below>
roles: 
  - id: <ID of the discord role>
    name: <short name of the role, used for logging>
    criteria: # list of criteria, evaluated with OR
      - conditions: # list of conditions, evaluated with AND
          - type: <see condition type options>
            value: <see condition type options>
            invert: (Optional - defaults to false) <true/false>
```

name_format_type options
```yaml
none: Do not adjust nicknames (will ignore title options, even if set)
first_last: First and Last Name (John Smith)
first_last_initial: First Name and Last Initial (John S)
cid: Certificate ID (1234567)
```

title_type options
```yaml
none: Do not append any title
division: Will use the same title formatting as the official server (i.e. ZZZ EC or ZZZ S3)
local: Similar to division, but omits the facility if it matches the server's facility.
rating: Short Rating (i.e. OBS or S3)
```

condition_type options
The "invert" option reverses the condition, so that the opposite is applied

| type             | value                              | description                                                                                   |
|------------------|------------------------------------|-----------------------------------------------------------------------------------------------|
| in_division      | true/false                         | Is the user in the VATUSA Division                                                            |
| division_visitor | true/false                         | Does the user visit any VATUSA subdivision                                                    |
| home             | facility ID                        | Is the user on the home roster of the specified facility                                      |
| visit            | facility ID                        | Is the user on the visitor roster of the specified facility                                   |
| home_or_visit    | facility ID                        | Is the user on the home or visitor roster of the specified facility                           |
| rating           | rating short (i.e. OBS, S3, I1)    | Does the user have the specified rating                                                       | 
| role             | role ID                            | Does the user have the specified role (at any facility)                                       |
| all              | n/a                                | True for any user who has a verified discord ID                                               |
| division_staff   | n/a                                | Is the user VATUSA Division Staff USA##                                                       |
| facility_staff   | facility ID                        | Does the user have a facility staff role at the specified facility (ATM, DATM, TA, EC, FE, WM |
| facility_role    | facility id : role (i.e "ZZZ:ATM") | Does the user have the specified role at the specified facility                               |
