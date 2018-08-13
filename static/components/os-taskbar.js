const taskbar = document.createElement('template');

taskbar.innerHTML = `
    <style>
    :host {
        z-index: 200;
        --taskbar-height: 48px;
    }
    nav {
        display: flex;
        flex-direction: row;
        height: var(--taskbar-height);
    }
    nav > div {
        display: flex;
        flex-direction: row;
        align-items: center;
    }
    .launcher-area {
        width: var(--taskbar-height);
    }
    .main-area {
        flex: 1;
        font-weight: 300;
        font-size: 0.8em;
        padding: 0 6px;
        border-bottom: 1px solid #A2C8FF;
    }
    .utilities-area {
        width: auto;
        border-bottom: 1px solid #A2C8FF;
    }
    .controls-area {
        padding: 0 6px;
        border-bottom: 1px solid #A2C8FF;
    }
    .controls-area button {
        background-color: #006AFF;
        color: #ffffff;
    }
    </style>
    <nav>
        <div class="launcher-area"><slot name="launcher-area">Launcher not implemented</slot></div>
        <div class="main-area"><slot name="main-area"></slot></div>
        <div class="utilities-area"><slot name="utilities-area"></slot></div>
        <div class="controls-area"><slot name="controls-area"></slot></div>
    </nav>
`;

customElements.define('os-taskbar', class extends HTMLElement {
    constructor() {
        super();
        let shadow = this.attachShadow({mode: 'open'});
        shadow.appendChild(taskbar.content.cloneNode(true))
    }
});