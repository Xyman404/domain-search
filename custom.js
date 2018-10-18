// Отправка заявки
$(document).ready(function() {
    $('#check-domain-form').submit(function() {
        if (document.form.value == "") {
            valid = false;
            return valid;
        }
        $.ajax({
            type: "POST",
            url: "http://127.0.0.1:3000/log",
            data: $(this).serialize()
        }).done(function(data) {
            console.log(data);
            $("#result").html("<p>" + data + "</p>");
            $(this).find('input').val('');
            $('#check-domain-form').trigger('reset');
        });
        return false;
    });
});
