# golang101-authenticate

# Authentication
การทำ authen ในปัจจุบันมีมากมายหลายแบบ ทุกครั้งที่จะทำระบบ authen ใหม่สำหรับผมแล้วมักจะนึกไม่ออกว่าควรจะทำด้วยอะไรดี โดยส่วนใหญ่ถ้านึกอะไร ไม่ออกก็จะหยิบ JWT มาใช้ตลอด เนื่องจากผมไม่ค่อยแม่นในการทำระบบ authen มากวันนี้เลยถือโอกาสลองหาข้อมูลศึกษา authen ในรูปแบบอื่นดูแล้วสรุปไว้อ่านเอง เพื่อครั้งหน้าจำไม่ได้จะได้เปิด content นี้มาดู เพื่อเตื่อนความจำ ถ้าผิดพลาดประการใด หรือผมสรุปอะไรผิดไป สามารถ comment บอกได้นะครับ 

โอเค เริ่มเลย 
ในบทความนี้ผมได้หาข้อมูลในการ authen มา 3 ตัว คือ HTTP Auth, API Key และ OAuth

## HTTP Auth
เป็นรูปแบบการทำ authen ที่ implement ได้ง่าย โดยหลักการของ HTTP Auth 
คือการนำ เอา username และ password มาทำการตรวจสอบว่าสามารถเข้าใช้งานในระบบเราได้ไหม ถ้าตรวจสอบแล้วว่าผ่าน 
จะนำเอา username และ password มาเข้ารหัสในรูปแบบ base64 แล้วนำ base64 ที่ได้มาใส่ไว้ในของ Http header ภายใต้ key Authorization 
โดย key ส่วนนี้จะถูกส่งไปในทุกๆ request ที่เรียกไป

และเมื่อผู้ใช่ได้รับการตรวจสอบสิทธ์ application จะร้องขอ password เสมอ แล้วผู้ใช้งานก็ไม่มีทางรู้ว่า application 
นำ password ของเราไปทำอะไรบ้าง หากมีคนอื่นที่รู้ password ของเรา ก็จะไม่สามารถรู้ได้ว่า มีใครนำ password ไปใช้เนื่องจาก basic authen ไม่ support การตรวจสอบสิทธ์แบบ multi factor

ดังนั้น HTTP Auth นั้นเหมาะสำหรับนำไปใช้การ authen ระหว่าง server to server ไม่เหมาะสำหรับ การทำระบบให้กับ Users

## API Key
สำหรับ API Key จะถูกใช้สำหรับการตรวจสอบสิทธ์ของ application ที่เขาถึง API โดยไม่ต้อง อ้างอิงจาก ผู้ใช้งานจริงๆ

โดย application จะต้องส่ง API Key มาในทุกๆ request ที่ส่งมาให้ API แล้ว API จะใช้ key ในการ ระบุตัว application และสิทธ์เข้าถึงได้

จะเห็นว่าการทำ authen ด้วย API Key ทำให้ฝั่ง client ใช้งานง่ายมาก เพียงแค่ client ส่ง API Key มาในทุกๆ request ก็สามารถเข้าใช้ API ได้แล้ว จะไม่เหมือน HTTP Auth ก่อนหน้า ที่ client ต้องกรอก username และ password เข้ามาด้วย

แต่ API Key นั้นจะถูกใช้สำหรับการ authen ระหว่าง application กับ application ไม่ใช่ผู้ใช้งานจริง และด้วยเหตุนี้ จึงเป็นเรื่องยากที่เราจะเก็บ API Key ไว้เป็นความลับ 

นอกจากนี้ API Key ก็ไม่ใช้การ authen ที่อยู่ในรูปแบบมาตรฐาน จึงมี API Key มากมายที่มีการ implement ที่แตกต่างกัน


# Reference 
https://nordicapis.com/the-difference-between-http-auth-api-keys-and-oauth/