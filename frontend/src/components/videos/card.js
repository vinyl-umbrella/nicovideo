import style from './style.css';
import util from '../../func/util';

const VideoCard = (props) => {
  function removePrefix(s) {
    return s.replace(/sm|nm|so|[^0-9]/g, '');
  }

  return (
    <div class={style.card}>
      <a href={`https://www.nicovideo.jp/watch_tmp/${props.id}`} target='_blank' rel='noopener noreferrer'>
        <img src={`https://nicovideo.cdn.nimg.jp/thumbnails/${removePrefix(props.id)}/${removePrefix(props.id)}`} />
        <div class={style.title}>{props.title}</div>
        <div class={style.info}>
          <div>{util.secToMinSec(props.duration)}</div>
          <div>
            {props.viewCount.toLocaleString('ja-JP')}視聴 &nbsp;
            {props.commentCount.toLocaleString('ja-JP')}コメ
          </div>
          <div>{util.formatDateTime(props.registeredAt)}</div>
        </div>
      </a>
    </div>
  );
};

export default VideoCard;
