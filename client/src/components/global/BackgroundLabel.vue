<template funcitonal>
   <div :class="['bg-label flex flex-center noselect', colorPreset]">
       {{ text }}
   </div>
</template>

<script>
export default {
name: 'BackgroundLabel',
    props: {
        colorPreset: {
            type: String,
            default: 'grey'
        },
        text: {
            type: String,
            default: 'Something went wrong'
        }
    }
}
</script>

<style lang='scss'>
.bg-label {
    padding: 2rem 1rem;
    border: 2px solid lighten($dark-grey, 40%);
    border-radius: .5rem;
    max-width: 20rem;
    opacity: .4;
    margin: 0 auto;
}
</style>