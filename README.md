http -v --json POST localhost:9000/api/v1/signup username=admin password=admin

http -v --json POST localhost:9000/api/v1/login username=admin password=admin

http -f GET localhost:9000/api/v1/hello "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjY4MTU5NTUsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTU2NjgxMjM1NX0.IVpVCxqZ8mDcCV5kIrsoGlQYIP7fAE3G4kNioTKa4t0"  "Content-Type: application/json"

http -f GET localhost:8000/auth/hello "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjY1Mzk5NzAsImlkIjoidGVzdCIsIm9yaWdfaWF0IjoxNTY2NTM2MzcwfQ.P_27MDxpjnFnBGgVopHpGeCxl0HUO5g4kZPRxKthcH8"  "Content-Type: application/json"

