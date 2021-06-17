### This script takes your Domeneshop API creds and creates .tf files for every dns and www records of your account
### Also create a .ps1 to initally import every records into your tfstate file (the extention can be changed to .sh on linux)


import requests
import json
import uuid

domainsID = requests.get("https://api.domeneshop.no/v0/domains", auth=("kcpS1etiquiKTh7w","NBwbCw7OeOTxDTLPk4z7ak2CGzDiqPewBxYkv4ESZhXubWzRrjn9KjFBpgW19olm"))
domains = json.loads(domainsID.text)
for i in range(len(domains)):
    domainID = str(domains[i]["id"])
    domainName = str(domains[i]["domain"])
    #print(domainID)
    
    print(domainName + ".tf")

    recordsID = requests.get("https://api.domeneshop.no/v0/domains/"+domainID+"/dns", auth=("kcpS1etiquiKTh7w","NBwbCw7OeOTxDTLPk4z7ak2CGzDiqPewBxYkv4ESZhXubWzRrjn9KjFBpgW19olm"))
    records = json.loads(recordsID.text)
    if len(records) == 0:
        print("yeet record : No DNS records")
    elif records == {'code': 'resource:unknown', 'help': 'See documentation at https://api.domeneshop.no/docs'}:
        print("yeet record : no DNS service")
    else:
        f = open(str("record-" + domainName + ".tf"), "w")
        g = open(str("terraimport.ps1"), "a")

        for j in range(len(records)):
            recordID = str(records[j]["id"])
            #print(domainID+"/"+recordID)
            domainName = str(domains[i]["domain"])
            recordHost = str(records[j]["host"])
            recordType = str(records[j]["type"])
            recordData = str(records[j]["data"]).replace("\"", "")
            recordTTL = str(records[j]["ttl"])
            domainUUID = str(uuid.uuid4().hex)
            domainTerraName = str("AUTOGEN-" + ((recordHost.replace('.', '-')).replace('@', 'apex')).replace('*', 'wildcard') + "-" + domainUUID)

#pls don't judge, it works :)

            print("resource \"domeneshop_record\" \"" + domainTerraName + "\" {")
            print("  domain_id = " + domainID)
            print("  host      = " + "\"" + recordHost + "\"")
            print("  type      = " + "\"" + recordType + "\"")
            print("  data      = " + "\"" + recordData + "\"")
            print("  ttl       = " + recordTTL)
            print("}")
            print("")

            f.write("resource \"domeneshop_record\" \"" + domainTerraName + "\" {" + '\n')
            f.write("  domain_id = " + domainID + '\n')
            f.write("  host      = " + "\"" + recordHost + "\"" + '\n')
            f.write("  type      = " + "\"" + recordType + "\"" + '\n')
            f.write("  data      = " + "\"" + recordData + "\"" + '\n')
            f.write("  ttl       = " + recordTTL + '\n')
            f.write("}" + '\n')
            f.write("" + '\n')

            g.write("terraform import domeneshop_record." + domainTerraName + " " + domainID + "/" + recordID + "\n")
        
        f.close()
        g.close()

    forwardsID = requests.get("https://api.domeneshop.no/v0/domains/"+domainID+"/forwards", auth=("kcpS1etiquiKTh7w","NBwbCw7OeOTxDTLPk4z7ak2CGzDiqPewBxYkv4ESZhXubWzRrjn9KjFBpgW19olm"))
    forwards = json.loads(forwardsID.text)
    if len(forwards) == 0:
        print("yeet forward : no forward record")
    elif records == {'code': 'resource:unknown', 'help': 'See documentation at https://api.domeneshop.no/docs'}:
        print("yeet forward : no forward service")
    else:
        f = open(str("forward-" + domainName + ".tf"), "w")
        g = open(str("terraimport.ps1"), "a")

        for k in range(len(forwards)):
            forwardHost = str(forwards[k]["host"])
            forwardFrame = str(forwards[k]["frame"])
            forwardUrl = str(forwards[k]["url"])
            forwardUUID = str(uuid.uuid4().hex)
            forwardTerraName = ("AUTOGEN-" + recordHost.replace('.', '-')).replace('@', 'apex').replace('*', 'wildcard') + "-" + domainName.replace('.', '-').replace('@', 'apex').replace('*', 'wildcard') + forwardUUID

            print("resource \"domeneshop_forward\" \"" + forwardTerraName + "\" {")
            print("  domain_id = " + domainID)
            print("  host      = " + "\"" + forwardHost + "\"")
            print("  url      = " + "\"" + forwardUrl + "\"")
            print("}")
            print("")

#pls don't judge, it works :)
            f.write("resource \"domeneshop_forward\" \"" + forwardTerraName + "\" {" + '\n')
            f.write("  domain_id = " + domainID + '\n')
            f.write("  host      = " + "\"" + forwardHost + "\"" + '\n')
            f.write("  url      = " + "\"" + forwardUrl + "\"" + '\n')
            f.write("}" + '\n')
            f.write("" + '\n')

            g.write("terraform import domeneshop_forward." + domainTerraName + " " + domainID + "/" + forwardHost + "\n")
        
        f.close()
        g.close()
