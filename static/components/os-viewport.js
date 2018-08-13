const viewport = document.createElement('template');

viewport.innerHTML = `
    <style>
    main {
        display: flex;
        flex-direction: column;
        height: 100vh;
        width: 100%;
    }
    .taskbar-area {
        height: 48px;
        background-color: #E7F1FF;
    }
    .workspace-area {
        display: flex;
        fex-direction: row;
        flex: 1;
    }
    .shortcuts {
        min-width: 48px;
        background-color: #E7F1FF;
        border-right: 1px solid #A2C8FF;
        box-sizing: border-box;
    }
    .workspace {
        display: flex;
        height: 100%;
    }

    </style>
    <main>
        <div class="taskbar-area"><slot name="taskbar-area"></slot></div>
        <div class="workspace-area">
            <div class="shortcuts"><slot name="shortcut-area"></slot></div>
            <div class="workspace"><slot name="workspace-area"></slot></div>
        </div>
    </main>
`;

customElements.define('os-viewport', class extends HTMLElement {
    constructor() {
        super();
        let shadow = this.attachShadow({mode: 'open'});
        shadow.appendChild(viewport.content.cloneNode(true))
    }
});