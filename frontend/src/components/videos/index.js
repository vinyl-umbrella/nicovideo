import style from './style.css';

import VideoCard from './card';

const Videos = (props) => {
  return (
    <div class={style.videos}>
      {props.v.map((v) => {
        return (
          <VideoCard
            key={v.vid}
            id={v.vid}
            title={v.title}
            duration={v.duration}
            viewCount={v.viewCount}
            commentCount={v.commentCount}
            registeredAt={v.registeredAt}
          />
        );
      })}
    </div>
  );
};

export default Videos;
