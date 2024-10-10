<div class="loading"></div>

<style>
  .loading {
    width: 40px;
    height: 40px;
    margin: 100px auto;
    background-color: #6a197d;

    border-radius: 100%;
    animation: scaleout 1s infinite ease-in-out;
  }

  @keyframes scaleout {
    0% {
      -webkit-transform: scale(0);
    }
    100% {
      -webkit-transform: scale(1);
      opacity: 0;
    }
  }
</style>
