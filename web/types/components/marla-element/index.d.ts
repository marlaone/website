import { TwLitElement } from "../../util/TwLitElement";
export declare class MyElement extends TwLitElement {
    render(): import("lit-html").TemplateResult<1>;
}
declare global {
    interface HTMLElementTagNameMap {
        'marla-element': MyElement;
    }
}
