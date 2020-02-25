# AIDBOX
## USE
### 1. Setup Aidbox.Dev https://docs.aidbox.app/installation/setup-aidbox.dev
### 2. Create .env file
```
# AIDBOX_
## insecure, (self sign self signed certificate)
AIDBOX_INSECURE=NO

AIDBOX_HOST=https://{{host}}.aidbox.app
AIDBOX_CLIENT={{client_name}}
AIDBOX_SECRET={{client_secret}}
```
### 3. Create first Patient

create_test.yml
```
resourceType: Patient
name:
- given: [Max1]
  family: Turikov1
gender: male
birthDate: '1990-10-10'
address:
- line:
  - 123 Oxygen St
  city: Hello1
```

```
cat create_test.yml | aidbox -ctype yml -create
```