import style from './style.css';

import { useEffect, useRef, useState } from 'preact/hooks';
import callApi from '../../func/callApi';
import Videos from '../videos';

const Home = () => {
  const searchText = useRef(null);
  const [videoList, setVideoList] = useState([]);

  useEffect(() => {
    if (searchText.current.value === '') {
      searchText.current.value = new URLSearchParams(window.location.search).get('q') || '';
      search();
    }
  }, []);

  async function search() {
    document.title = `「${searchText.current.value}」の検索結果 - ニコ動 (Re:仮) 検索サービス (仮)`;
    if (searchText.current.value == '') return;
    const j = await callApi.search(searchText.current.value, 1);
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
    await search();
  }

  return (
    <div class={style.home}>
      <form onSubmit={onSubmit} class={style.searchbox}>
        <input type='text' placeholder='Search...' ref={searchText} />
        <button type='submit'>検索</button>
      </form>
      {searchText != '' && videoList.length === 0 ? <div>検索結果 0件</div> : <Videos v={videoList} />}
    </div>
  );
};

export default Home;
