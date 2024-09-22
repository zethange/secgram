import { Component, JSX } from 'solid-js';

interface TabButtonProps {
  children?: JSX.Element,
  isActive: boolean,
  onClick: () => void,
}

export const TabButton: Component<TabButtonProps> = (props) => {
  return (
    <button
      class={'rounded-lg p-1 w-full text-center transition hover:-translate-y-0.5 hover:scale-105 ' + (props.isActive ? 'bg-black text-white' : 'bg-white border')} onClick={props.onClick}>
      {props.children}
    </button>
  );
}