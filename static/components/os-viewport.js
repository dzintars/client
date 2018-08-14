const viewport = document.createElement('template');

viewport.innerHTML = `
    <style>
    .viewport {
        display: flex;
        flex-direction: column;
        height: 100vh;
        width: 100%;
    }
    .taskbar-area {
        height: 48px;
        width: 100%;
        background-color: #E7F1FF;
    }
    .workspace-area {
        display: flex;
        fex-direction: row;
        flex: 1;
        width: 100%;
    }
    .shortcuts {
        width: 48px;
        min-width: 48px;
        background-color: #E7F1FF;
        box-sizing: border-box;
        border-right: 1px solid #A2C8FF;
    }
    .workspace {
        display: flex;
        flex: 1;
        //border: 3px solid #A2C8FF;
    }

    </style>
    <div class="viewport">
        <div class="taskbar-area"><slot name="os-taskbar"></slot></div>
        <div class="workspace-area">
            <div class="shortcuts"><slot name="os-shortcut"></slot></div>
            <div class="workspace"><slot name="os-workspace"></slot></div>
        </div>
    </div>
`;

customElements.define('os-viewport', class extends HTMLElement {
    constructor() {
        super();
        let shadow = this.attachShadow({mode: 'open'});
        shadow.appendChild(viewport.content.cloneNode(true))
    }
});