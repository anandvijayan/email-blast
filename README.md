# Personalized email blast utility
Create a personalized email and send a blast to all recipients specified in the config file.

## Install tools
Install go lang and nodejs

`make install-tools`

## Format the Code
Format the code

`make format`


## Email Config
Update the email config with the details of recipients and email message

[email_blast.yaml](./config/email_blast.yaml)

```
EmailBlast:
    Recepients:
        - Name: Anand Vijayan
          Email: anand_vijayan@yahoo.com
    Message:
        SalutationPrefix: Dear
        MessageSubject: 'PLEASE CONFIRM! District 2 Gavel Club area contest'
        MessageTone: Formal
        MessageContent: "The District 2 Gavel clubs are having an area contest and are in need of judges. 
                         There are 2 sets of speaking and evaluation contests based on participants' age. 
                         If you cannot be there for the whole five hours, you have an option to judge just 1 set of contests. Food and drinks shall be served.\n\n
                         10:00 am - 2:00 pm, Saturday 7th June 2025\n
                         PARK RIDGE CHURCH\n
                         3805 Malty Rd, Bothell, WA 98012\n\n
                         Please let me know your availability."
        Signature: "---\n
                    Anand Vijayan (alias Andy)\n
                    Youth Leadership Chair\n
                    #: PN-06645899\n
                    anand.vijayachandran@gmail.com\n
                    602-738-6626\n
                    https://d2tm.org/youth-programs-gavel-clubs/"

```

## Environment Variables
Set environment variables

```
export OPENAI_API_KEY=<the openai api key for enhanced message>
export EMAIL_ID=<email id of user who is sending the blast message>
export EMAIL_PASSWORD=<password of the user sending the email>
export ENHANCE_MESSAGE=<Yes/No flag indicating if the message needs to be enhanced with AI>
```

## Test and Send Email Notifications
Test the enhanced message generation

`make test`

Use [email_blast.yaml](./config/email_blast.yaml) config and send personalized email notifications.

`make send-notification`
