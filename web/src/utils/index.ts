import { FlowComponent, JSX } from "solid-js"

export function getChildren(el: Element): Element[] {
  const contents = []
  for(let i = 0; i <el.childNodes.length;i++) {
    contents.push(el.childNodes[i])
  }
  return contents
}

export function findElement(elementName: string, cb: (el: Element) => void): void {
  const els = document.getElementsByTagName(elementName)
  for(let i = 0; i < els.length; i++) {
    cb(els[i])
  }
}