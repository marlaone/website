import { html } from 'lit';
import {TwLitElement} from "../../util/TwLitElement";

export class MyElement extends TwLitElement {
    render() {
        return html`
          <button class="btn"><slot></slot></button>
        `;
    }
}

window.customElements.define('marla-element', MyElement);

declare global {
    interface HTMLElementTagNameMap {
        'marla-element': MyElement
    }
}
