const viewport = document.createElement('template');

viewport.innerHTML = `
    <style>
    main {
        display: flex;
        flex-direction: column;
        height: 100vh;
        width: 100%;
    }
    .taskbar {
        height: 48px;
    }
    .workspace {
        display: flex;
        height: 100%;
    }

    </style>
    <main>
        <div class="taskbar"><slot name="taskbar-area"></slot></div>
        <div class="workspace"><slot name="workspace-area"></slot></div>
    </main>
`;

customElements.define('os-viewport', class extends HTMLElement {
    constructor() {
        super();
        let shadow = this.attachShadow({mode: 'open'});
        shadow.appendChild(viewport.content.cloneNode(true))
    }
});