function fadeIn(element) {
    // ToDO: Need to hide other profiles if visible.
    var target = document.getElementById(element);
    target.style.display = "block";
}

function fadeOut(element) {
    var target = document.getElementById(element);
    target.style.display = "none";
}

let table = document.getElementById("tbl-delivery-orders"), rIndex;
for (var i= 0; i < table.rows.length; i++) {
    table.rows[i].ondblclick = function() {
        rIndex = this.rowIndex;
        fadeIn("order-profile");
        form = document.getElementById("delivery-order-profile-form");

        form.elements[0].value = this.cells[1].firstChild.innerHTML;
        form.elements[1].value = this.cells[0].firstChild.innerHTML;
        form.elements[2].value = this.cells[2].innerHTML;
        form.elements[3].value = this.cells[7].innerHTML;
        form.elements[4].value = this.cells[3].innerHTML;
        form.elements[5].value = this.cells[4].innerHTML;
        form.elements[6].value = this.cells[5].innerHTML;
        form.elements[7].value = this.cells[6].innerHTML;
    }
}

function toggle(widget, slot) {
    'use strict';

    let template = document.getElementById(widget);
    let templateContent = template.content;
    
    // Attempt to reference the element in the document, not the template content
    var imported = template.firstChild;

    // Check for the element, not the template content
    if (document.body.contains(imported)) {
        // Element exists, get its parent
        imported.parentNode.removeChild(imported);
    } else {
        // Use .importNode to bring template content in:
        //document.querySelectorAll('[slot="ws-main-panel"]');
        document.body.appendChild(document.importNode(templateContent, true));
    }
}