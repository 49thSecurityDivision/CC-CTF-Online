
var apiUrl = "http://backend.carolinacon.kctf.cloud:4242";

function getMenuItems() {
  $.ajax({
    url: apiUrl + "/menu",
    type: "GET",
    beforeSend: function (xhr) { xhr.setRequestHeader('X-Authenticate', localStorage.getItem("the-pig-authentication")) },
    success: function (result) {
      menuHtml = "";
      for (var item in result) {
        menuHtml += '<div class="menu-item">';
        menuHtml += result[item].item;
        menuHtml += '<span>$';
        menuHtml += result[item].price;
        menuHtml += '</span></div>';
      }
      $(".menu-list").html(menuHtml);
    },
    error: function (result) {
      console.error(result);
      $(".menu-list").html('<div class="menu-item menu-error">There was an error fetching the menu 😥</div>');
    }
  });
}

$(document).ready(function () {

  getMenuItems();

});