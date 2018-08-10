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
        box-shadow: 0 3px 4px 0 rgba(0,0,0,0.2), 0 3px 3px -2px rgba(0,0,0,0.14), 0 1px 8px 0 rgba(0,0,0,0.12);
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
    }
    .utilities-area {
        width: auto;
    }
    .controls-area {
        padding: 0 6px;
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