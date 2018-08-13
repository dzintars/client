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
        border-top: 1px solid #A2C8FF;
    }
    .sidebar {
        display: none;
    }
    .main-panel {
        flex: 1;
    }
    .right-panel {
        flex: 1;
    }
    .statusbar {
        display: flex;
        flex-direction: row;
        align-items: center;
        height: 22px;
        border-top: 1px solid #e6e6e6;
        font-size: 0.7em;
        padding: 0 6px;
    }

    </style>
    <div class="workspace">
        <nav class="toolbar"><slot name="toolbar"></slot></nav>
        <div class="panel-area">
            <div class="sidebar"><slot name="sidebar"></slot></div>
            <div class="main-panel"><slot name="main-panel"></slot></div>
            <div class="right-panel"><slot name="right-panel"></slot></div>
        </div>
        <div class="statusbar"><slot name="statusbar"></slot></div>
    </div>
`;

customElements.define('os-workspace', class extends HTMLElement {
    constructor() {
        super();
        let shadow = this.attachShadow({mode: 'open'});
        shadow.appendChild(workspace.content.cloneNode(true))
    }
});