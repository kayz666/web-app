$(document).ready(function(){
    // $.validator.setDefaults({
    //     submitHandler: function() {
    //
    //     }
    // });

    $("#signinWithEmailForm").validate({
        submitHandler:function () {
            form.submit();
        }
    });

    $("#registerWithEmailForm").validate({
        submitHandler:function () {
            form.submit();
        },
        rules:{
            email:{
                required:true,
                email:true
            },
            password:{
                required:true,
                minlength:5
            },
            confirm_password:{
                required:true,
                minlength:5,
                equalTo: "#password"
            }
        },
        message:{
            password:{
                required:"please provide a password",
                minlength:"Your password must be at least 5 characters long"

            },
            confirm_password:{
                required:"please provide a password",
                minlength:"Your password must be at least 5 characters long",
                equalTo: "Please enter the same password as above"
            },
            email:"please enter a valid email address"
        }
    });

});