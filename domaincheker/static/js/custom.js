$(document).ready(function() {
    $('#check-domain-form').submit(function() {
        $("#loader").show("slow");
        $.ajax({
            url:     "http://127.0.0.1:3000/log", //url страницы
            type:     "POST", //метод отправки
            data: $("#input-domain").serialize(),  // Сеарилизуем объект
        }).done(function(data){
            console.log(data);
            $("#result").html(data);
            $(this).find('input').val('');
            $('#check-domain-form').trigger('reset');
            $("#loader").hide("slow");
        });
        return false;
    });
});
