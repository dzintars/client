function fadeIn(element) {
    var target = document.getElementById(element);
    target.style.display = "block";
}

function fadeOut(element) {
    var target = document.getElementById(element);
    target.style.display = "none";
}

function openProfile() {
    // Get selected row
    let table = document.getElementById("tbl-delivery-orders"),rIndex;
    for (var i= 0; i < table.rows.length; i++) {
        table.rows[i].onclick = function() {
            rIndex = this.rowIndex;
            console.log(rIndex);
        }
    }
}

let table = document.getElementById("tbl-delivery-orders"),rIndex;
for (var i= 0; i < table.rows.length; i++) {
    table.rows[i].ondblclick = function() {
        rIndex = this.rowIndex;
        fadeIn("order-profile")
        document.forms[11].elements[0].value = this.cells[1].firstChild.innerHTML;
        document.forms[11].elements[1].value = this.cells[0].firstChild.innerHTML;
        document.forms[11].elements[2].value = this.cells[2].innerHTML;
        document.forms[11].elements[3].value = this.cells[7].innerHTML;
        document.forms[11].elements[4].value = this.cells[3].innerHTML;
        document.forms[11].elements[5].value = this.cells[4].innerHTML;
        document.forms[11].elements[6].value = this.cells[5].innerHTML;
        document.forms[11].elements[7].value = this.cells[6].innerHTML;
    }
}