import style from './style.css';

import { useEffect, useRef, useState } from 'preact/hooks';
import callApi from '../../func/callApi';
import Page from './page';
import Videos from '../videos';

const Home = () => {
  const searchText = useRef('');
  const [videoList, setVideoList] = useState([]);

  useEffect(() => {
    searchText.current.value = new URLSearchParams(window.location.search).get('q') || '';
    const page = new URLSearchParams(window.location.search).get('page') || 1;

    search(page);
  }, []);

  async function search(page) {
    if (searchText.current.value === '') {
      document.title = 'ニコ動 (Re:仮) 検索サービス (仮)';
    } else {
      document.title = `「${searchText.current.value}」の検索結果 - ニコ動 (Re:仮) 検索サービス (仮)`;
    }

    const j = await callApi.search(searchText.current.value, page);
    setVideoList([...j]);
  }

  function updateHistory() {
    const url = new URL(window.location.href);
    const searchParams = url.searchParams;
    searchParams.set('q', searchText.current.value);
    url.search = searchParams.toString();
    const newUrl = url.toString();

    console.log(newUrl);
    window.history.pushState({ searchText: searchText.current.value }, '', newUrl);
  }

  async function onSubmit(e) {
    e.preventDefault();
    updateHistory();
    await search(1);
  }

  return (
    <div class={style.home}>
      <form onSubmit={onSubmit} class={style.searchbox}>
        <input type='text' placeholder='Search...' ref={searchText} />
        <button type='submit'>検索</button>
      </form>
      {videoList.length === 0 ? <div>検索結果 0件</div> : <Videos v={videoList} />}
      <Page />
    </div>
  );
};

export default Home;
