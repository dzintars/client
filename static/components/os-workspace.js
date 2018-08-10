const workspace = document.createElement('template');

workspace.innerHTML = `
    <style>
    :host {
        height: 100%;
        width: 100%;
    }
    .workspace {
        display: flex;
        flex-direction: column;
        height: 100%;
    }
    .toolbar {
        display: flex;
        flex-direction: row;
        align-items: center;
        height: 32px;
    }

    .panel-area {
        display: flex;
        flex: 1;
        flex-direction: row;
    }
    .sidebar {
        width: 20px;
    }
    .main-panel {
        flex: 1;
    }
    .right-panel {
        flex: 1;
    }

    </style>
    <div class="workspace">
        <nav class="toolbar"><slot name="toolbar"></slot></nav>
        <div class="panel-area">
            <div class="sidebar"><slot name="sidebar"></slot></div>
            <div class="main-panel"><slot name="main-panel"></slot></div>
            <div class="right-panel"><slot name="right-panel"></slot></div>
        </div>
    </div>
`;

customElements.define('os-workspace', class extends HTMLElement {
    constructor() {
        super();
        let shadow = this.attachShadow({mode: 'open'});
        shadow.appendChild(workspace.content.cloneNode(true))
    }
});