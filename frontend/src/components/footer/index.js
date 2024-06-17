import style from './style.css';

const Footer = () => (
  <footer class={style.footer}>
    <div>
      当サイトによって発生したいかなる損害についても，当方は一切の責任を負いません．ただ，ニコニコ動画を楽しむためのサービスを提供することを目的としています．
    </div>
    <div>
      <a href='https://github.com/vinyl-umbrella/nicovideo' target='_blank' rel='noopener noreferrer'>
        GitHub
      </a>
    </div>
  </footer>
);

export default Footer;
