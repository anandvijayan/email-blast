# Toastmasters International Club DCP Status Checker
Check the Club DCP status page to generate the current DCP report in Yaml and image formats.

## Install tools
Install go lang and nodejs

`make install-tools`

## Format the Code
Format the code

`make format`

## Email Config
Update the email config with the details of recipients and email message

![dcp_report.yaml](./config/email_blast.yaml)

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


Capture the Club DCP report screenthot png

`make generate-screenshot`

![Club DCP Report](./reports/dcp_report.png)

Generate Club DCP report in Yaml and capture the Club DCP report screenthot

`make generate-all`

Use the generated Yaml report and summarize the achievements using Chat GPT API. Send an announcement message via Email with the summary and image captured.

`make send-notification`
