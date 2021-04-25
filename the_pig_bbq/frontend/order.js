
var apiUrl = "http://backend.carolinacon.kctf.cloud:4242";


function getMenuItems() {
    $.ajax({
        url: apiUrl + "/menu",
        type: "GET",
        beforeSend: function (xhr) { xhr.setRequestHeader('X-Authenticate', localStorage.getItem("the-pig-authentication")) },
        success: function (result) {
            optionsHtml = "";
            for (var item in result) {
                $("#orderDropdown").append(new Option(result[item].item, result[item].item))
            }
        },
        error: function (result) {
            console.error(result);
            $(".order-form").html('<div class="menu-item menu-error">There was an error fetching the menu items 😥</div>');
        }
    });
}

function getFlag() {
    $("body").html("<div style='text-align: center; padding: 50px'>" + atob("ZmxhZ3tyVlMxMWctWEJnV2FoN0hUWld6Zm1Vdm9NY25tSXZEc30=") + "</div>");
}

function submitOrder() {
    var order = $("#orderDropdown").val();
    $.ajax({
        url: apiUrl + "/order",
        type: "POST",
        data: { "order": order },
        success: function (result) {
            $(".order-form").html('<div class="menu-item menu-error">Your order was submitted! Order number: ' + result.orderNumber + '</div>');
        },
        error: function (result) {
            console.error(result);
            $(".order-form").html('<div class="menu-item menu-error">There was an error submitting your order 😥</div>');
        }
    });
}

$(document).ready(function () {

    getMenuItems();

    $("#orderSubmit").on('click touchstart', function () {
        submitOrder();
    });

});