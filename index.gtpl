<!DOCTYPE html>
<html lang="en" >
<head>
  <meta charset="UTF-8">
  <title>CSS Validation</title>
  
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/normalize/5.0.0/normalize.min.css">

  
      <link rel="stylesheet" href="css/style.css">

  
</head>

<body>
  <form action="/register" method="post" class="form">
	<div class="form__field">
		<input type="text" name="surname" placeholder="Фамилия*" required />
	</div>
	<div class="form__field">
		<input type="text" name="firstname" placeholder="Имя*" required />
	</div>
	<div class="form__field">
		<input type="text" name="endname" placeholder="Отчество*" required />
	</div>
	<div class="form__field">
		<input type="email" name="email" placeholder="E-Mail*"  required/>
		<span class="form__error">Это поле должно содержать E-Mail в формате example@site.com</span>
	</div>
	<div class="form__field">
		<input type="tel" name="name" placeholder="Телефон*" pattern="[\+]\d{1}\s[\(]\d{3}[\)]\s\d{3}[\-]\d{2}[\-]\d{2}" minlength="18" maxlength="18" />
		<span class="form__error">Это поле должно содержать телефон в формате +7 (123) 456-78-90</span>
	</div>
	<div class="form__field">
		<input type="text" name="univerkafedra" placeholder="Ваша кафедра*" required/>
	</div>
	<div class="form__field">
		<input type="text" name="univergroup" placeholder="Ваша группа*" required />
	</div>
	<div class="form__field">
		<input type="checkbox" required name="terms"> Я согласен на обработку данных
	</div>
	
	<button type="submit">Отправить</button>
</form>
  
  
</body>
</html>
