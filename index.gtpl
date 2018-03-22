<!DOCTYPE html>
<html lang="en" >
<head>
  <meta charset="UTF-8">
  <title>Форма регистрации</title>
  
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/normalize/5.0.0/normalize.min.css">

    <style type="text/css">
    body {
    	text-align: center;
    	padding: 10px;
    }

    .form {
    	display: inline-block;
    	padding: 25px;
    	border: 1px solid #ccc;
    	border-radius: 5px;
    	width: 500px;
    }

    .form__field {
    	margin-bottom: 10px;
    }

    .form__error {
    	color: red;
    	text-align: left;
    	font-size: 12px;
    	display: block;
    	margin-top: 3px;
    	display: none;
    }

    .form input {
    	outline: none;
    	display: block;
    	width: 100%;
    	border-radius: 2px;
    	border: 1px solid #ccc;
    	padding: 10px;
    	box-sizing: border-box;
    }

    .form button {
    	width: 100%;
    	padding: 10px;
    	border-radius: 2px;
    	border: 0;
    	background-color: #ccc;
    	color: #fff;
    }

    input:valid:not(:placeholder-shown) {
    	border-color: green;
    }

    input:invalid:not(:placeholder-shown) {
    	border-color: red;
    }
    input:invalid:not(:placeholder-shown) + .form__error {
    	display: block;
    }
    </style>


  
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
		<input type="tel" name="tel" placeholder="Телефон*" pattern="[\+]\d{1}\s[\(]\d{3}[\)]\s\d{3}[\-]\d{2}[\-]\d{2}" minlength="18" maxlength="18" />
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
