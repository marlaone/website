import { FlowComponent, JSX } from 'solid-js'

export type MarlaElement<P = {}> = FlowComponent<P, Element[]|JSX.Element>

export * from './alert'