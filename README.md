# devflowv2

bu proje go diliyle yazılmış basit bir command line uygulamasıdır.  
amaç user, project ve task gibi şeyleri oluşturmak, listelemek, güncellemek ve silmek.  
her şey bellekte (memory'de) tutuluyor. database yok.

bu projeyi go öğrenmek için step-by-step geliştirdim.

## proje yapısı

cmd/  
→ main.go burada. uygulama buradan başlatılıyor.

models/  
→ user, project ve task struct'ları burada tanımlı  
→ ayrıca onları tuttuğumuz map'ler de burada

handlers/  
→ user_handler.go: user işlemleri  
→ project_handler.go: project işlemleri  
→ task_handler.go: task işlemleri  
→ hepsi create, list, update, delete içeriyor

## ilk versiyon (v0.1 diyebilirim)

başta map[string]string kullanarak user'ları tuttum  
örneğin id -> name gibi  
crud işlemlerini handlers/user_handler.go içinde yazdım  
main.go dosyasından çağırarak test ettim

## v0.2'de yaptıklarım

go'da struct yapısını öğrendim  
user'ı artık struct olarak tanımladım: id, name, email

map[string]User yapısına geçtik  
crud fonksiyonları struct'a göre güncellendi  
aynı şekilde project ve task için de struct'lar oluşturduk

task struct'ı içinde ProjectID alanı var  
bu sayede bir görevin hangi projeye ait olduğunu biliyoruz  
task'ı oluştururken geçerli bir proje id verilmesi gerekiyor

## dikkat ettiğim bazı şeyler

- update işlemlerinde struct'ı map'ten alıp alanlarını değiştirdim, sonra geri map'e koydum
- map erişiminde `value, ok := map[key]` şeklinde kontrol yaptım
- ikinci değer (ok) daima bool olur. true = key var, false = yok

## çalıştırmak için

önce şu komutu verdim:
go mod init github.com/kursatartar/devflowv2

sonra uygulamayı çalıştırmak için:
go run cmd/main.go

## örnek senaryo

handlers.CreateUser("1", "kürşat", "k@k.com")  
handlers.CreateProject("p1", "devflow", "go projesi")  
handlers.CreateTask("t1", "giriş fonksiyonu yaz", "todo", "p1")

handlers.ListUsers()  
handlers.ListProjects()  
handlers.ListTasks()

handlers.UpdateUser(...)  
handlers.DeleteProject(...)  
vb.

## diğer notlar

- repo'da her değişiklikten sonra git add, commit ve push yaptım
- bellekte veri tuttuğum için uygulama kapanınca her şey silinir. kalıcılık yok.
- user, project ve task’lar arasında şu an yalnızca task → project ilişkisi var
- tüm id’ler elle veriliyor. otomatik id üretimi yapılmadı
- aynı id ile tekrar create yaparsam eskisinin üstüne yazar, uyarı verilmez  
- input kullanıcıdan alınmıyor, main.go'dan manuel test yapılıyor
