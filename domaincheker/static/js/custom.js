// Отправка заявки
$(document).ready(function() {
    $('#check-domain-form').submit(function() {
        if (document.form.name.value == "") {
            valid = false;
            return valid;
        }
        $.ajax({
            type: "POST",
            url: "http://127.0.0.1:3000/log",
            data: $(this)
        }).done(function(data) {
            console.log(data);
            $("#result").text(data);
            $(this).find('input').val('');
            $('#check-domain-form').trigger('reset');
        });
        return false;
    });
});
